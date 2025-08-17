// This stub file contains items that aren't used yet; feel free to remove this module attribute
// to enable stricter warnings.
#![allow(unused)]

/// various log levels
#[derive(Clone, PartialEq, Debug)]
pub enum LogLevel {
    Info,
    Warning,
    Error,
}
/// primary function for emitting logs
pub fn log(level: LogLevel, message: &str) -> String {
    use LogLevel::*;
    
    (match level {
        Info => info,
        Warning => warn,
        Error => error,
    })(message)
}
pub fn info(message: &str) -> String {
    format!("[INFO]: {message}").to_string()
}
pub fn warn(message: &str) -> String {
    format!("[WARNING]: {message}").to_string()
}
pub fn error(message: &str) -> String {
    format!("[ERROR]: {message}").to_string()
}
