import TambahPenyakit from "./pages/TambahPenyakit";
import Search from "./pages/Search";
import TesDNA from "./pages/TesDNA";
import { Routes, Route } from "react-router-dom";
import Navbar from "./pages/NavBar";
import "./pages/Style.css";
function App() {
  return (
    <div className="App">
      <Navbar />
      <Routes>
        <Route path="/" element={<Search />} />
        <Route path="/TambahPenyakit" element={<TambahPenyakit />} />
        <Route path="/TesDNA" element={<TesDNA />} />
      </Routes>
    </div>
  );
}
/*tes*/
export default App;
