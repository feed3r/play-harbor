use std::env;
use std::process::Command;
use std::thread::sleep;
use std::time::Duration;
use sysinfo::{Processes, System, SystemExt, PidExt};

fn main() {
    let args: Vec<String> = env::args().collect();

    if args.len() != 3 {
        eprintln!("ERROR: Needs launch URL and EXE Name");
        eprintln!("Usage: epic_game_launcher <epicUrl> <exeName>");
        return;
    }

    let epic_url = &args[1];
    let exe_name = &args[2];

    println!("Starting url: {}", epic_url);

    // Lancia l’URL tramite Windows (equivalente a "start <url>")
    if let Err(e) = Command::new("rundll32")
        .arg("url.dll,FileProtocolHandler")
        .arg(epic_url)
        .spawn()
    {
        eprintln!("Failed to start URL: {}", e);
        return;
    }

    // Aspetta un po’ che il gioco parta
    sleep(Duration::from_secs(5));

    // Carica la lista dei processi
    let mut sys = System::new_all();
    sys.refresh_processes();

    let mut game_pid = None;

    for (pid, process) in sys.processes() {
        let name = process.name().to_lowercase();
        if name == exe_name.to_lowercase() || name == format!("{}.exe", exe_name.to_lowercase()) {
            game_pid = Some(pid.clone());
            break;
        }
    }

    if game_pid.is_none() {
        eprintln!("Could not find a process with name: {}", exe_name);
        return;
    }

    println!("Game started. Waiting for it to exit...");

    // Loop finché il processo esiste
    loop {
        sys.refresh_processes();
        if !sys.processes().contains_key(&game_pid.unwrap()) {
            break;
        }
        sleep(Duration::from_secs(2));
    }

    println!("Game exited.");
}
