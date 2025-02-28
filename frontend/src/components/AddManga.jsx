import React, { useState } from 'react';

const AddManga = () => {
  const [formData, setFormData] = useState({
    title: '',
    series: '',
    volume: '',
    cover: '',
    path: '',
  });

  const handleSubmit = (e) => {
    e.preventDefault();
    fetch('/api/manga/add', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(formData),
    })
      .then((response) => {
        if (response.ok) {
          alert('Manga added successfully!');
          setFormData({ title: '', series: '', volume: '', cover: '', path: '' }); // Reset form
        } else {
          alert('Failed to add manga. Please try again.');
        }
      })
      .catch((error) => console.error('Error adding manga:', error));
  };

  return (
    <form onSubmit={handleSubmit} className="bg-white p-6 rounded-lg shadow-md">
      <h2 className="text-xl font-semibold mb-4">Add Manga</h2>
      <div className="space-y-4">
        <input
          type="text"
          placeholder="Title"
          value={formData.title}
          onChange={(e) => setFormData({ ...formData, title: e.target.value })}
          className="w-full p-2 border rounded"
          required
        />
        <input
          type="text"
          placeholder="Series"
          value={formData.series}
          onChange={(e) => setFormData({ ...formData, series: e.target.value })}
          className="w-full p-2 border rounded"
          required
        />
        <input
          type="number"
          placeholder="Volume"
          value={formData.volume}
          onChange={(e) => setFormData({ ...formData, volume: e.target.value })}
          className="w-full p-2 border rounded"
          required
        />
        <input
          type="text"
          placeholder="Cover URL"
          value={formData.cover}
          onChange={(e) => setFormData({ ...formData, cover: e.target.value })}
          className="w-full p-2 border rounded"
          required
        />
        <input
          type="text"
          placeholder="File Path"
          value={formData.path}
          onChange={(e) => setFormData({ ...formData, path: e.target.value })}
          className="w-full p-2 border rounded"
          required
        />
        <button
          type="submit"
          className="w-full bg-blue-500 text-white p-2 rounded hover:bg-blue-600"
        >
          Add Manga
        </button>
      </div>
    </form>
  );
};

export default AddManga; 