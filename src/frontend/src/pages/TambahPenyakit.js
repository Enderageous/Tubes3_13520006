function TambahPenyakit() {
  return (
    <div>
      <h1 className="subtitle"> Tambah Penyakit</h1>
      <div className="tambahPenyakit">
        <div className="twocolumn">
          <p>Nama Pengguna:</p>
          <form>
            <input
              className="twoInput"
              type="text"
              name="penyakit"
              placeholder="penyakit..."
            />
          </form>
        </div>
        <div className="twocolumn">
          <p>Sequence DNA:</p>
          <button
            className="uploadButtonPenyakit"
            variant="contained"
            color="primary"
            component="span"
          >
            Upload
          </button>
        </div>
        <div className="onecolumn">
          <button className="submitButton" type="submit">
            Submit
          </button>
        </div>
      </div>
    </div>
  );
}

export default TambahPenyakit;
