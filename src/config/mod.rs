use serde::Deserialize;
use std::error::Error;
use std::fs;

#[derive(Debug, Deserialize)]
pub struct Config {
    pub bin_base_dir: String,
}

impl Config {
    pub fn from_yaml_string(s: &str) -> Result<Self, serde_yaml::Error> {
        serde_yaml::from_str(s)
    }
}

pub fn load_from_yaml_file(path: &str) -> Result<Config, Box<dyn Error>> {
    let s = fs::read_to_string(path)?;

    let c = Config::from_yaml_string(&s)?;
    Ok(c)
}
