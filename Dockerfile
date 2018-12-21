# Stage #1
# The base image of builder contains Dep, Go Dependancy Management tool
FROM lushdigital/docker-golang-dep as builder

# Set the path where the source code will reside in the container
ARG src_path=/go/src/github.com/elgarni/hala-art
ADD . $src_path
WORKDIR $src_path

# Ensure all dependancies are present in vendor
RUN dep ensure --vendor-only

# Run the tests. If the tests fail, then the build will be interrupted
RUN go test

# Build the go app binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app main.go

# Stage #2
# The base image will be Scratch
FROM scratch

# Copy the go binary from stage #1
COPY --from=builder app /

# Set the entry point to be the binary
ENTRYPOINT ["/app"]