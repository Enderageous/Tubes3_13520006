import { useState, useEffect } from "react";
import axios from "axios";

function TesDNA() {
  const [diseases, setDiseases] = useState([]);

  useEffect(() => {
    axios
      .get(`https://enigmatic-brook-59106.herokuapp.com/api/disease`)
      .then((res) => {
        console.log(res.data)
        setDiseases(res.data);
      });
  }, []);

  return (
    <div className="body">
      <h1 className="subtitle">Tes DNA</h1>
      <div className="tesDNA">
        <div className="threecolumn">
          <p>Nama Pengguna:</p>
          <form>
            <input
              className="threeInput"
              type="text"
              name="nama"
              placeholder="<pengguna>"
            />
          </form>
        </div>
        <div className="threecolumn">
          <p>Sequence DNA:</p>
          <form>
            <button
              className="uploadButtonDNA"
              variant="contained"
              color="primary"
              component="span"
            >
              Upload
            </button>
            <br></br>
            <br />
            <button className="submitButton1" type="submit">
              Submit
            </button>
          </form>
        </div>
        <div className="threecolumn">
          <p>Prediksi Penyakit:</p>
          <form>
            {/* <input
              className="threeInput"
              type="text"
              name="nama"
              placeholder="<penyakit>"
            /> */}
            <select>
              {diseases.map((disease) => (
                <option key={disease.id} value={disease.disease_id}>
                  {disease.disease_name}
                </option>
              ))}
            </select>
          </form>
        </div>
      </div>
    </div>
  );
}

export default TesDNA;
