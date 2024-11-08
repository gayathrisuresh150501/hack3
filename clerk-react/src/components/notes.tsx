import React from 'react';

function Note({ notes }) {
  return (
    <div>
      {notes.length === 0 ? (
        <p>No notes found. Create one!</p>
      ) : (
        <ul>
          {notes.map(note => (
            <li key={note._id}>
              <h3>{note.title}</h3>
              <p>{note.content}</p>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}

export default Note;
