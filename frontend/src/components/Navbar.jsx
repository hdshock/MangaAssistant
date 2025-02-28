import React from 'react';
import { Link } from 'react-router-dom';

const Navbar = () => {
  return (
    <nav className="bg-white shadow-md p-4">
      <div className="container mx-auto flex justify-between items-center">
        <h1 className="text-xl font-semibold">Manga Assistant</h1>
        <div className="space-x-4">
          <Link to="/check-duplicates" className="text-gray-700 hover:text-blue-500">Check Duplicates</Link>
        </div>
      </div>
    </nav>
  );
};

export default Navbar; 