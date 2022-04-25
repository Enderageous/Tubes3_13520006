function Search() {
  return (
    <div>
      <h1 className="subtitle"> Search</h1>
      <div className="search">
        <form>
          <input
            className="searchInput"
            type="text"
            name="penyakit"
            placeholder="penyakit..."
          />
        </form>
      </div>
    </div>
  );
}

export default Search;
