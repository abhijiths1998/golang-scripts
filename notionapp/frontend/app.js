let currentPage = null;

async function fetchPages() {
  const res = await fetch('http://localhost:8081/api/pages');
  const pages = await res.json();
  const list = document.getElementById('pageList');
  list.innerHTML = '';
  pages.forEach(p => {
    const li = document.createElement('li');
    const a = document.createElement('a');
    a.href = '#';
    a.textContent = p.title;
    a.onclick = () => loadPage(p.id);
    li.appendChild(a);
    list.appendChild(li);
  });
}

async function createPage() {
  const title = document.getElementById('newPageTitle').value;
  const res = await fetch('http://localhost:8081/api/pages', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ title })
  });
  await res.json();
  document.getElementById('newPageTitle').value = '';
  fetchPages();
}

async function loadPage(id) {
  const res = await fetch(`http://localhost:8081/api/pages/${id}`);
  const page = await res.json();
  currentPage = page.id;
  document.getElementById('pageTitle').textContent = page.title;
  document.getElementById('markdown').value = page.content || '';
  renderTodos(page.todos);
  renderSubPages(page.children);
}

function renderTodos(todos) {
  const list = document.getElementById('todoList');
  list.innerHTML = '';
  todos.forEach(t => {
    const li = document.createElement('li');
    const cb = document.createElement('input');
    cb.type = 'checkbox';
    cb.checked = t.done;
    cb.onchange = () => updateTodo(t.id, cb.checked);
    li.appendChild(cb);
    li.appendChild(document.createTextNode(' '+t.text));
    list.appendChild(li);
  });
}

async function addTodo() {
  const text = document.getElementById('newTodoText').value;
  await fetch(`http://localhost:8081/api/pages/${currentPage}`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ text })
  });
  document.getElementById('newTodoText').value = '';
  loadPage(currentPage);
}

async function updateTodo(id, done) {
  await fetch(`http://localhost:8081/api/pages/${currentPage}/todos/${id}`, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ done })
  });
}

function renderSubPages(ids) {
  const list = document.getElementById('subPageList');
  list.innerHTML = '';
  ids.forEach(async id => {
    const res = await fetch(`http://localhost:8081/api/pages/${id}`);
    const page = await res.json();
    const li = document.createElement('li');
    const a = document.createElement('a');
    a.href = '#';
    a.textContent = page.title;
    a.onclick = () => loadPage(page.id);
    li.appendChild(a);
    list.appendChild(li);
  });
}

async function createSubPage() {
  const title = document.getElementById('newSubPageTitle').value;
  await fetch('http://localhost:8081/api/pages', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ title, parent_id: currentPage })
  });
  document.getElementById('newSubPageTitle').value = '';
  loadPage(currentPage);
  fetchPages();
}

fetchPages();
