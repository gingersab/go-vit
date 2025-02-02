use std::process::{Command, Stdio};

// Learn more about Tauri commands at https://tauri.app/develop/calling-rust/
#[tauri::command]
fn start_go_server() {
    Command::new("../../go-vit")
        .stdout(Stdio::null())
        .stderr(Stdio::null())
        .spawn()
        .expect("Failed to start Go WebSocket server");
}

#[cfg_attr(mobile, tauri::mobile_entry_point)]
pub fn run() {
    tauri::Builder::default()
        .plugin(tauri_plugin_opener::init())
        .invoke_handler(tauri::generate_handler![start_go_server])
        .run(tauri::generate_context!())
        .expect("error while running tauri application");
}
