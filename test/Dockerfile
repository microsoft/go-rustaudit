FROM rust:1.62.1-bullseye as builder
RUN git clone https://github.com/rust-secure-code/cargo-auditable.git
WORKDIR /cargo-auditable/cargo-auditable
RUN cargo build
WORKDIR /cargo-auditable/cargo-auditable/tests/fixtures/workspace
RUN /cargo-auditable/target/debug/cargo-auditable auditable build
FROM scratch
COPY --from=builder \
  /cargo-auditable/cargo-auditable/tests/fixtures/workspace/target/debug/crate_with_features_bin \
  /crate_with_features_bin
