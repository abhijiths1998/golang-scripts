import Database from 'better-sqlite3';
import path from 'path';
import fs from 'fs';

const dbPath = path.join(process.cwd(), 'data.db');
const firstRun = !fs.existsSync(dbPath);
const db = new Database(dbPath);

if (firstRun) {
  db.exec(`
    CREATE TABLE pages (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      title TEXT
    );
    CREATE TABLE todos (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      page_id INTEGER,
      text TEXT,
      done INTEGER
    );
  `);
}

export default db;
