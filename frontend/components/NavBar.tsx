import React from 'react';
import Link from 'next/link';

const NavBar: React.FC = () => {
  return (
    <nav className="bg-blue-500 p-4 text-white">
      <div className="container mx-auto flex justify-between">
        <Link href="/" className="text-xl font-bold">
          Blog
        </Link>
        <div>
          <Link href="/register" className="mr-4">
            Register
          </Link>
          <Link href="/login">
            Login
          </Link>
        </div>
      </div>
    </nav>
  );
};

export default NavBar;
