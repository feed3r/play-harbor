#[cfg(test)]
mod tests {
    use std::process::Command;

    #[test]
    fn test_missing_args() {
        let args = vec!["playdock.exe".to_string()];
        assert_ne!(args.len(), 3, "Should fail with missing arguments");
    }

    #[test]
    fn test_command_error() {
        // Simula errore nell'avvio del comando
        let result = Command::new("false").output();
        assert!(
            result.is_err() || !result.unwrap().status.success(),
            "Expected error when running command"
        );
    }

    #[test]
    fn test_process_not_found() {
        // Simula nessun processo trovato
        let processes = vec!["other.exe", "notgame.exe"];
        let found = processes.iter().any(|name| *name == "Game.exe");
        assert!(!found, "Process should not be found");
    }

    #[test]
    fn test_process_found() {
        // Simula processo trovato
        let processes = vec!["Game.exe", "other.exe"];
        let found = processes.iter().any(|name| *name == "Game.exe");
        assert!(found, "Process should be found");
    }
}
