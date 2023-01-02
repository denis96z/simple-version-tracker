use clap::Parser;

mod config;

#[derive(Parser, Debug)]
#[clap(author, version, about, long_about = None)]
struct Args {
    #[clap(short, long)]
    config: String,
}

fn main() {
    let args = Args::parse();

    let c = config::load_from_yaml_file(&args.config).unwrap();
    dbg!(c);
}
