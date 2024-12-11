# Use a compatible Go version
FROM golang:1.23.2 as builder

# Install necessary tools
RUN go install github.com/gogo/protobuf/protoc-gen-gogofaster@latest && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# Install and set up libbls
ENV libbls https://github.com/ibalajiarun/sgxwallet/releases/download/v0.4/libbls.tar.gz
RUN wget -O /root/libbls.tar.gz ${libbls} && \
    mkdir -p /opt/libbls && \
    tar -C /opt/libbls -xzf /root/libbls.tar.gz && \
    rm /root/libbls.tar.gz

# Install and set up libpolycrypto
ENV libpolycrypto https://github.com/ibalajiarun/libpolycrypto/releases/download/v0.2/libpolycrypto.tar.gz
RUN wget -O /root/libpolycrypto.tar.gz ${libpolycrypto} && \
    mkdir -p /opt/libpolycrypto && \
    tar -C /opt/libpolycrypto -xzf /root/libpolycrypto.tar.gz && \
    rm /root/libpolycrypto.tar.gz

# Set up the working directory and copy source code
WORKDIR /app
COPY . .

# Set SGX environment variables
ENV SGX_SDK=/opt/sgxsdk
ENV PATH=$PATH:$SGX_SDK/bin:$SGX_SDK/bin/x64
ENV PKG_CONFIG_PATH=$PKG_CONFIG_PATH:$SGX_SDK/pkgconfig
ENV LD_LIBRARY_PATH=$LD_LIBRARY_PATH:$SGX_SDK/sdk_libs

# Set Go environment variables
ENV CGO_ENABLED=1 GOOS=linux GOARCH=amd64
ENV LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/app/utils/sgx/enclave_trinc:/app/utils/sgx/enclave_trudep
ENV SGX_MODE=HW

# Build the project
RUN make build

###############################################################################

# Runtime image
FROM balajia/sgx-runtime:latest

LABEL org.opencontainers.image.source https://github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher

# Add Intel SGX repository GPG key
RUN wget -qO - https://download.01.org/intel-sgx/sgx_repo/ubuntu/intel-sgx-deb.key | gpg --dearmor -o /usr/share/keyrings/intel-sgx-archive-keyring.gpg

# Add the Intel SGX repository
RUN echo "deb [arch=amd64 signed-by=/usr/share/keyrings/intel-sgx-archive-keyring.gpg] https://download.01.org/intel-sgx/sgx_repo/ubuntu bionic main" > /etc/apt/sources.list.d/intel-sgx.list

# Update and install necessary dependencies
RUN apt-get update && \
    apt-get install -y libgomp1 && \
    rm -rf /var/lib/apt/lists/* && \
    rm -rf /var/cache/apt/archives/*

# Copy built binaries and libraries from the builder
# COPY --from=builder /app/enclaves/trinx/enclave_trinx/libtrinx.so /lib/libtrinx.so
# COPY --from=builder /app/enclaves/trinx/enclave_trinx/trinxenclave.signed.so /lib/trinxenclave.signed.so
# COPY --from=builder /app/enclaves/threshsign/enclave_threshsign/libthreshsign.so /lib/libthreshsign.so
# COPY --from=builder /app/enclaves/threshsign/enclave_threshsign/threshsignenclave.signed.so /lib/threshsignenclave.signed.so
# COPY --from=builder /app/enclaves/usig/enclave_usig/libusig.so /lib/libusig.so
# COPY --from=builder /app/enclaves/usig/enclave_usig/usig_enclave.signed.so /lib/usig_enclave.signed.so
# COPY --from=builder /app/enclaves/threshsign2/enclave_threshsign/libthreshsign2.so /lib/libthreshsign2.so
# COPY --from=builder /app/enclaves/threshsign2/enclave_threshsign/threshsign2enclave.signed.so /lib/threshsign2enclave.signed.so
COPY --from=builder /app/bin/peer /bin/peer
COPY --from=builder /app/bin/client /bin/client

# Expose the required port
EXPOSE 7000

# Set the entry point
ENTRYPOINT ["/bin/peer"]