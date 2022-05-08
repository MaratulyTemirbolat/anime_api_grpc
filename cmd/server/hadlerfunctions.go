// package main

// import (
// 	"google.golang.org/grpc/codes"
// 	"google.golang.org/grpc/status"
// )

// const zeroValue int = 0

// func is_id_provided(userId uint32) error {
// 	if userId == uint32(zeroValue) {
// 		return status.Errorf(codes.InvalidArgument, "You can not send ID = 0")
// 	}
// 	return nil
// }
