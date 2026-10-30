
import { useEffect, useState } from "react";

type Task = {

  id: number;
  description: string;
  completed: boolean;
};

export default function App() {

  const [tasks, setTasks] = useState<Task[]>([]);
  const [description, setDescription] = useState("");
  const [error, setError] = useState("");

  async function loadTasks() {
    try {

      const response = await fetch("/api/tasks");

      if (!response.ok) throw new Error("Cannot load tasks");

      setTasks(await response.json());

      setError("");
    } catch {
        
      setError("Cannot connect to the Go API.");
    }
  }

  useEffect(() => {

    void loadTasks();
  }, []);

  async function addTask() {

    if (!description.trim()) return;

    try {

      const response = await fetch("/api/tasks", {

        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ description })
      });

      if (!response.ok) throw new Error("Cannot add task");

      setDescription("");
      await loadTasks();

    } catch {

      setError("Could not add the task.");
    }
  }

  async function toggleTask(task: Task) {
    try {

      const response = await fetch(`/api/tasks/${task.id}`, {
        method: "PATCH",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ completed: !task.completed })
      });

      if (!response.ok) throw new Error("Cannot update task");

      await loadTasks();
    } catch {

      setError("Could not update the task.");
    }
  }

  async function deleteTask(id: number) {
    try {

      const response = await fetch(`/api/tasks/${id}`, {

        method: "DELETE"
      });

      if (!response.ok) throw new Error("Cannot delete task");

      await loadTasks();
    } catch {

      setError("Could not delete the task.");
    }
  }

  return (
    <main style={{

      maxWidth: 650,
      margin: "60px auto",
      padding: 30,
      fontFamily: "Arial, sans-serif"
    }}>
      <h1>My To-Do List</h1>
      <p>React + Go + PostgreSQL + Docker</p>

      <div style={{ display: "flex", gap: 10 }}>
        <input
          value={description}

          //onChange Handles typing with setDescription:
          //updates the text stored in React state as you type
          onChange={
            
            e => setDescription(e.target.value)
        }
          onKeyDown={e => {
            
            if (e.key === "Enter") void addTask();
          }}

          placeholder="Enter a task..."
          style={{ flex: 1, padding: 12, fontSize: 16 }}
        />

        <button onClick={() => void addTask()}>
          Add Task
        </button>
      </div>

      {error && <p style={{ color: "red" }}>{error}</p>}

      <ul style={{ listStyle: "none", padding: 0 }}>
        {tasks.map(task => (
          <li key={task.id} style={{
            display: "flex",
            alignItems: "center",
            gap: 12,
            padding: 12,
            borderBottom: "1px solid #ddd"
          }}>
            <input
              type="checkbox"
              checked={task.completed}
              onChange={() => void toggleTask(task)}
            />

            <span style={{
              flex: 1,
              textDecoration: task.completed ? "line-through" : "none"
            }}>
              {task.description}
            </span>

            <button onClick={() => void deleteTask(task.id)}>
              Delete
            </button>
          </li>
        ))}
      </ul>
    </main>
  );
}
