### Client & server example of gRPC

For create pb files
1. cd protos
2. mkdir -p gen/go
2. run command (install protoc before it)

protoc -I proto proto/music/music.proto --go_out=./gen/go/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative


    Вместо пути до конкретного proto-файла можете использовать вот такую запись:
    -I proto proto/sso/*.proto <прочие параметры>
    Тогда будет сгенерирован код по всем proto-файлам из указанной директории (не сработает для родной командной строки Windows). Рекуррентно заходить в поддиректории оно в таком виде не будет.