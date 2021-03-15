PSM="ad.fe.demos"

run:
	go run main.go

build:
	go build -o output/bin/${PSM}
	# exec output/bin/${PSM}