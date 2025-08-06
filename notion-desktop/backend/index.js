import express from 'express';
import cors from 'cors';
import db from '../database/db.js';

const app = express();
app.use(cors());
app.use(express.json());

// Pages CRUD
app.get('/api/pages', (req, res) => {
  const pages = db.prepare('SELECT id, title FROM pages').all();
  res.json(pages);
});

app.post('/api/pages', (req, res) => {
  const { title } = req.body;
  const result = db.prepare('INSERT INTO pages (title) VALUES (?)').run(title);
  res.json({ id: result.lastInsertRowid, title });
});

app.put('/api/pages/:id', (req, res) => {
  const { title } = req.body;
  db.prepare('UPDATE pages SET title = ? WHERE id = ?').run(title, req.params.id);
  res.json({ id: req.params.id, title });
});

app.delete('/api/pages/:id', (req, res) => {
  db.prepare('DELETE FROM pages WHERE id = ?').run(req.params.id);
  res.json({});
});

// Simple todos
app.get('/api/todos', (req, res) => {
  const todos = db.prepare('SELECT id, page_id, text, done FROM todos').all();
  res.json(todos);
});

app.post('/api/todos', (req, res) => {
  const { page_id, text } = req.body;
  const result = db.prepare('INSERT INTO todos (page_id, text, done) VALUES (?, ?, 0)').run(page_id, text);
  res.json({ id: result.lastInsertRowid, page_id, text, done: 0 });
});

app.put('/api/todos/:id', (req, res) => {
  const { text, done } = req.body;
  db.prepare('UPDATE todos SET text = ?, done = ? WHERE id = ?').run(text, done, req.params.id);
  res.json({ id: req.params.id, text, done });
});

app.delete('/api/todos/:id', (req, res) => {
  db.prepare('DELETE FROM todos WHERE id = ?').run(req.params.id);
  res.json({});
});

const port = process.env.PORT || 3001;
app.listen(port, () => {
  console.log(`Backend running on port ${port}`);
});

export default app;
