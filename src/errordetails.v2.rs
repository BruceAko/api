// This file is @generated by prost-build.
/// Backend is error detail for Backend.
#[derive(serde::Serialize, serde::Deserialize)]
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Backend {
    /// Backend error message.
    #[prost(string, tag = "1")]
    pub message: ::prost::alloc::string::String,
    /// Backend HTTP response header.
    #[prost(map = "string, string", tag = "2")]
    pub header: ::std::collections::HashMap<
        ::prost::alloc::string::String,
        ::prost::alloc::string::String,
    >,
    /// Backend HTTP status code.
    #[prost(int32, optional, tag = "3")]
    pub status_code: ::core::option::Option<i32>,
}
/// Unknown is error detail for Unknown.
#[derive(serde::Serialize, serde::Deserialize)]
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Unknown {
    /// Unknown error message.
    #[prost(string, optional, tag = "1")]
    pub message: ::core::option::Option<::prost::alloc::string::String>,
}
