import { useEffect, useState } from 'react';
import './App.css';
import { addNewList, getAllLists } from './service/api';

function App() {
  const [lists, setLists] = useState([]);
  useEffect(() => {
    getAllLists()
      .then(setLists)
      .catch(console.error);
  }, []);

  const [title, setTitle] = useState("");

  async function addList() {
    if (!title) return;

    try {
      const newList = await addNewList({title});
      setLists([...lists, newList]);
      setTitle("");
    } catch (err) {
      console.log(err);
    }
  }

  return (
    <div className="App">
      <ul>
        {lists.map(l => (
          <li key={l.id}>{l.title}</li>
        ))}
      </ul>
      <p>
        <input
          value={title}
          onInput={e => setTitle(e.target.value)}
        />
        <button onClick={addList}>+</button>
      </p>
    </div>
  );
}

export default App;
