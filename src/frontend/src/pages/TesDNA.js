import { useState, useEffect } from "react";
import axios from "axios";
import { Button } from "@mui/material";

function TesDNA() {
  const [diseases, setDiseases] = useState([]);
  const [fileChoosen, setFileChoosen] = useState(false);
  const [currFile, setCurrFile] = useState();

  const onChange = (e) => {
    if (e.target.files[0]) {
      setFileChoosen(true);
      setCurrFile(e.target.files[0].name);
    } else {
      setFileChoosen(false);
    }
  };

  useEffect(() => {
    axios
      .get(`https://enigmatic-brook-59106.herokuapp.com/api/disease`)
      .then((res) => {
        console.log(res.data);
        setDiseases(res.data);
      })
      .catch((error) => console.log(error));
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
          <Button
            variant="contained"
            component="label"
            style={{
              borderRadius: 5,
              backgroundColor: "white",
              fontSize: "20px",
              color: "grey",
            }}
          >
            {fileChoosen ? `${currFile}` : "Upload file"}
            <input
              id="dnaSequence"
              type="file"
              name="sequence_dna"
              onChange={onChange}
              hidden
            ></input>
          </Button>
          <br></br>
          <br />
          <button className="submitButton1" type="submit" color="green">
            Submit
          </button>
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
