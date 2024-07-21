FROM rust:1.79.0-bullseye as builder
RUN git clone https://github.com/rust-secure-code/cargo-auditable.git
WORKDIR /cargo-auditable/cargo-auditable
RUN cargo build
RUN rustup target add wasm32-unknown-unknown
WORKDIR /cargo-auditable/cargo-auditable/tests/fixtures/wasm_crate
RUN /cargo-auditable/target/debug/cargo-auditable auditable build --target wasm32-unknown-unknown
WORKDIR /cargo-auditable/cargo-auditable/tests/fixtures/workspace
RUN /cargo-auditable/target/debug/cargo-auditable auditable build
FROM scratch
COPY --from=builder \
  /cargo-auditable/cargo-auditable/tests/fixtures/wasm_crate/target/wasm32-unknown-unknown/debug/wasm_crate.wasm \
  /cargo-auditable/cargo-auditable/tests/fixtures/workspace/target/debug/crate_with_features_bin \
  /
