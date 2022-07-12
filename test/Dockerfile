FROM rust:1.62.0-bullseye as builder
RUN git clone https://github.com/Shnatsel/rust-audit.git
RUN mkdir /output && cd rust-audit && cargo build && cp target/debug/hello-auditable /output
FROM scratch
COPY --from=builder /output /
