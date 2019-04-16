package main

import (
		"context"
			"log"
				"net"
					"sync"

						pb "github.com/naichadouban/gomicroservices/consignment-service/proto/consignment"
							"google.golang.org/grpc"
								"google.golang.org/grpc/reflection"
							)

							const (
									port = ":8010"
								)

								type repository interface {
										Create(*pb.Consignment) (*pb.Consignment, error)
											GetAll() []*pb.Consignment
										}
										type Repository struct {
												mu           sync.Mutex
													consignments []*pb.Consignment
												}

												// create a new consignment
												func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
														repo.mu.Lock()
															defer repo.mu.Unlock()
																repo.consignments = append(repo.consignments, consignment)
																	return consignment, nil
																}
																func (repo *Repository) GetAll() []*pb.Consignment {
																		return repo.consignments
																	}

																	// servie should implement all method to satify the service we defined in our protobuf definition.
																	// you can check the interface in the generated code itself for exact method signatures etc
																	// to give you a better idea.
																	type service struct {
																			repo *Repository
																		}

																		func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {
																				consignment, err := s.repo.Create(req)
																					if err != nil {
																								return nil, err
																									}
																										return &pb.Response{
																													Consignment: consignment,
																															Created:     true,
																																}, nil
																															}

																															func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
																																	consignments := s.repo.GetAll()
																																		return &pb.Response{
																																					Consignments: consignments,
																																						}, nil
																																					}

																																					func main() {
																																							repo := &Repository{}
																																								// set-up our grpc server
																																									lis, err := net.Listen("tcp", port)
																																										if err != nil {
																																													log.Panicf("Failed to listen:%v\n", err)
																																														}
																																															s := grpc.NewServer()
																																																pb.RegisterShippingServiceServer(s, &service{repo})
																																																	// register reflection service on grpc server
																																																		reflection.Register(s)
																																																			log.Println("grpc server listen on", port)
																																																				if err := s.Serve(lis); err != nil {
																																																							log.Panicf("grpc server error:%v\n", err)
																																																								}
																																																							}

