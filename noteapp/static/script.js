async function fetchNotes() {
    const res = await fetch('/notes');
    const notes = await res.json();
    const list = document.getElementById('notes');
    list.innerHTML = '';
    notes.forEach(n => {
        const li = document.createElement('li');
        li.textContent = n.text;
        list.appendChild(li);
    });
}

async function addNote() {
    const text = document.getElementById('noteText').value;
    await fetch('/notes', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ text })
    });
    document.getElementById('noteText').value = '';
    fetchNotes();
}

window.onload = fetchNotes;
