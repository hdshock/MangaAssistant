import React, { useState } from 'react';

const CheckDuplicates = () => {
  const [duplicates, setDuplicates] = useState([]);

  const handleCheckDuplicates = () => {
    fetch('/api/manga/check-duplicates')
      .then((response) => response.json())
      .then((data) => setDuplicates(data))
      .catch((error) => console.error('Error checking duplicates:', error));
  };

  const handleScan = async () => {
    try {
      const response = await fetch('http://localhost:2828/scan', {
        method: 'POST',
      });
      if (response.ok) {
        alert('Scan initiated successfully!');
      } else {
        alert('Failed to initiate scan.');
      }
    } catch (error) {
      console.error('Error:', error);
      alert('An error occurred while initiating the scan.');
    }
  };

  return (
    <div className="bg-white p-6 rounded-lg shadow-md">
      <h2 className="text-xl font-semibold mb-4">Check for Duplicates</h2>
      <button
        onClick={handleCheckDuplicates}
        className="bg-blue-500 text-white p-2 rounded hover:bg-blue-600"
      >
        Check Duplicates
      </button>
      <button
        onClick={handleScan}
        className="bg-blue-500 text-white px-4 py-2 rounded ml-2"
      >
        Scan Library
      </button>
      {duplicates.length > 0 && (
        <div className="mt-4">
          <h3 className="text-lg font-semibold">Duplicate Manga:</h3>
          <ul className="space-y-2">
            {duplicates.map((manga, index) => (
              <li key={index} className="p-2 border rounded">
                <p><strong>Title:</strong> {manga.Title}</p>
                <p><strong>Series:</strong> {manga.Series}</p>
                <p><strong>Volume:</strong> {manga.Volume}</p>
              </li>
            ))}
          </ul>
        </div>
      )}
    </div>
  );
};

export default CheckDuplicates; 