protoc greet/greetpb/greetpb.proto --go_out=plugins=grpc:.
protoc hw_sum/sumpb/sum.proto --go_out=plugins=grpc:.
protoc streaming_greet/GreetStreampb/greetStream.proto --go_out=plugins=grpc:.
protoc hw_streaming_number_decomposition/NumberDecompositionpb/decomposition.proto --go_out=plugins=grpc:.
protoc hw_client_streaming/avrgpb/avrg.proto --go_out=plugins=grpc:.
protoc hw_bidirectional_streaming/maxpb/max.proto --go_out=plugins=grpc:.
