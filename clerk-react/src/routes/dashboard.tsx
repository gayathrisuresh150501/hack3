import React, { useState, useEffect } from "react";
// import  Note from "../components/notes"
import {
  useUser,
  OrganizationList,
  OrganizationProfile,
  OrganizationSwitcher,
  CreateOrganization,
} from "@clerk/clerk-react";

const notes = [
  {
    _id: "1",
    title: "Note 1",
    content: "This is the first note.",
  },
  {
    _id: "2",
    title: "Note 2",
    content: "This is the second note.",
  },
];

const backendUrl="http://localhost:8000"

export default function DashboardPage() {
  const user = useUser();
  const [notes, setNotes] = useState([]);
  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");
  const [isEditing, setIsEditing] = useState(null);
  const [editTitle, setEditTitle] = useState("");
  const [editContent, setEditContent] = useState("");
  const [organizations, setOrganizations] = useState<[]>([]);

  const [plan, setPlan] = useState("free");

  useEffect(() => {
    // Fetch organization memberships when component mounts or user ID changes
    if (user.user?.getOrganizationMemberships) {
      user.user
        .getOrganizationMemberships()
        .then((orgs) => {
          setOrganizations(orgs.data || []); // Set organizations, handle empty
        })
        .catch((error) => {
          console.error("Error fetching organization memberships:", error);
        });
    }
  }, [user.user?.id, organizations.length]);

  useEffect(() => {
    fetchNotes();
  }, []);

  const fetchNotes = async () => {
    try {
      const response = await fetch(`${backendUrl}/api/notes`);
      const data = await response.json();
      setNotes(data); 
    } catch (error) {
      console.error("Error fetching notes:", error);
    }
  };

  // Create a new note by sending it to the backend
  const handleCreateNote = async (e) => {
    e.preventDefault();
    try {
      const response = await fetch(`${backendUrl}/api/notes`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify({ title, content }),
      });
      const newNote = await response.json();
      setNotes([...notes, newNote]);
      setTitle("");
      setContent("");
    } catch (error) {
      console.error("Error creating note:", error);
    }
  };

  // Edit a note
  const handleEditNote = (note) => {
    setIsEditing(note._id);
    setEditTitle(note.title);
    setEditContent(note.content);
  };

  const handleSaveEdit = async (id) => {
    try {
      const response = await fetch(`${backendUrl}/api/notes/${id}`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify({ title: editTitle, content: editContent }),
      });
      const updatedNote = await response.json();
      setNotes(
        notes.map((note) => (note._id === id ? updatedNote : note))
      );
      setIsEditing(null);
      setEditTitle("");
      setEditContent("");
    } catch (error) {
      console.error("Error updating note:", error);
    }
  };

  // Delete a note by ID
  const handleDeleteNote = async (id) => {
    try {
      await fetch(`${backendUrl}/api/notes/${id}`, {
        method: "DELETE",
      });
      setNotes(notes.filter((note) => note._id !== id));
    } catch (error) {
      console.error("Error deleting note:", error);
    }
  };

  return (
    <div >
    <div style={{display:"flex",alignItems:"center",justifyContent:"center",minHeight:"100vh",flexDirection:"column",minWidth:"80vw"}}>
      <h1>Dashboard page</h1>
      <p>This is a protected page.</p>

      <p>Plan: {plan}</p>
      <select onChange={(e) => setPlan(e.target.value)} value={plan}>
        <option value="free">Free Plan</option>
        <option value="pro">Pro Plan</option>
        <option value="enterprise">Enterprise</option>
      </select>
      <p>
        Welcome back, {user.user?.fullName} and userID: {user.user?.id}!
      </p>
      <p>
        Organizations that are present :{" "}
        {organizations?.map((org) => JSON.stringify(org.organization.name))}
      </p>

      {/* <div><OrganizationProfile path="/organization-profile" /></div> */}
      <div>
        <OrganizationSwitcher />
      </div>
      <div>
        <CreateOrganization path="/create-organization" />
      </div>
      <p>Here are your notes:</p>

      {/* Note List */}
      {notes.map((note) => (
        <div
          key={note.id}
          style={{
            border: "1px solid #ccc",
            padding: "10px",
            margin: "10px 0",
          }}
        >
          {isEditing === note.id ? (
            <>
              <input
                type="text"
                value={editTitle}
                onChange={(e) => setEditTitle(e.target.value)}
              />
              <textarea
                value={editContent}
                onChange={(e) => setEditContent(e.target.value)}
              />
              <button onClick={() => handleSaveEdit(note.id)}>Save</button>
            </>
          ) : (
            <>
              <h3>{note.title}</h3>
              <p>{note.content}</p>
              <button onClick={() => handleEditNote(note)}>Edit</button>
            </>
          )}
          <button onClick={() => handleDeleteNote(note.id)}>Delete</button>
          <button onClick={() => handleShareNote(note)}>Share</button>
        </div>
      ))}

      <form onSubmit={handleCreateNote}>
        <input
          type="text"
          placeholder="Title"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
        />
        <textarea
          placeholder="Content"
          value={content}
          onChange={(e) => setContent(e.target.value)}
        />
        <button type="submit">Create note</button>
      </form>
    </div>
    </div>
  );
}
