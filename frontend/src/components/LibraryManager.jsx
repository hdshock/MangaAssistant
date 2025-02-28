import React, { useState, useEffect } from 'react';

const LibraryManager = () => {
  const [libraries, setLibraries] = useState([]);
  const [newLibraryPath, setNewLibraryPath] = useState('');

  useEffect(() => {
    fetchLibraries();
  }, []);

  const fetchLibraries = () => {
    fetch('http://localhost:2828/api/libraries')
      .then((response) => response.json())
      .then((data) => setLibraries(data))
      .catch((error) => console.error('Error fetching libraries:', error));
  };

  const addLibrary = () => {
    fetch('http://localhost:2828/api/libraries/add', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ path: newLibraryPath }),
    })
      .then((response) => {
        if (response.ok) {
          fetchLibraries();
          setNewLibraryPath('');
        } else {
          alert('Failed to add library. Please check the path and try again.');
        }
      })
      .catch((error) => console.error('Error adding library:', error));
  };

  const removeLibrary = (path) => {
    fetch('http://localhost:2828/api/libraries/remove', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ path }),
    })
      .then((response) => {
        if (response.ok) {
          fetchLibraries();
        } else {
          alert('Failed to remove library. Please try again.');
        }
      })
      .catch((error) => console.error('Error removing library:', error));
  };

  return (
    <div className="bg-white p-6 rounded-lg shadow-md">
      <h2 className="text-xl font-semibold mb-4">Library Manager</h2>
      <div className="space-y-4">
        <div className="flex space-x-2">
          <input
            type="text"
            placeholder="Enter library path"
            value={newLibraryPath}
            onChange={(e) => setNewLibraryPath(e.target.value)}
            className="w-full p-2 border rounded"
          />
          <button
            onClick={addLibrary}
            className="bg-blue-500 text-white p-2 rounded hover:bg-blue-600"
          >
            Add Library
          </button>
        </div>
        <div>
          <h3 className="text-lg font-semibold mb-2">Watched Libraries</h3>
          {libraries.length > 0 ? (
            <ul className="space-y-2">
              {libraries.map((library, index) => (
                <li key={index} className="flex justify-between items-center p-2 border rounded">
                  <span>{library.path}</span>
                  <button
                    onClick={() => removeLibrary(library.path)}
                    className="bg-red-500 text-white p-1 rounded hover:bg-red-600"
                  >
                    Remove
                  </button>
                </li>
              ))}
            </ul>
          ) : (
            <p>No libraries being watched.</p>
          )}
        </div>
      </div>
    </div>
  );
};

export default LibraryManager; 