import React from "react";
import { Link } from "react-router-dom";
import "./Navbar.css";
function Navbar() {
  return (
    <div className="navbar">
      <ul>
        <li>
          <Link to="/">Search</Link>
        </li>
        <li>
          <Link to="/TambahPenyakit">TambahPenyakit</Link>
        </li>
        <li>
          <Link to="/TesDNA">TesDNA</Link>
        </li>
      </ul>
    </div>
  );
}

export default Navbar;
