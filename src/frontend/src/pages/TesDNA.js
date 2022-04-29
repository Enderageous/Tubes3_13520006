import { useState, useEffect } from "react";
import axios from "axios";
import { Button } from "@mui/material";

const TesDNA = () => {
  const [diseases, setDiseases] = useState([]);
  const [patientName, setPatientName] = useState("");
  const [diseaseId, setDiseasesId] = useState("");
  const [fileChoosen, setFileChoosen] = useState(false);
  const [dnaFiles, setDnaFiles] = useState(null);

  const onChange = (e) => {
    if (e.target.files[0]) {
      setFileChoosen(true);
      setDnaFiles(e.target.files[0]);
    } else {
      setFileChoosen(false);
    }
  };

  const handleSubmit = (event) => {
    event.preventDefault();
    const formData = new FormData();
    formData.append("disease_id", diseaseId);
    formData.append("patient_name", patientName);
    formData.append("dna_file", dnaFiles);

    axios({
      url: "https://enigmatic-brook-59106.herokuapp.com/api/prediction",
      data: formData,
      headers: { "Content-Type": "multipart/form-data" },
      method: "post",
    })
      .then((res) => alert("Data prediksi berhasil ditambahkan"))
      .catch((error) => console.log(error));
  };
  useEffect(() => {
    axios
      .get(`https://enigmatic-brook-59106.herokuapp.com/api/disease`)
      .then((res) => {
        setDiseases(res.data);
      })
      .catch((error) => console.log(error));
  }, []);

  return (
    <div className="body">
      <form onSubmit={handleSubmit}>
        <h1 className="subtitle">Tes DNA</h1>
        <div className="tesDNA">
          <div className="threecolumn">
            <p>Nama Pengguna:</p>
            <input
              id="nama"
              className="threeInput"
              type="text"
              name="patientName"
              onChange={(e) => setPatientName(e.target.value)}
              placeholder="<pengguna>"
            />
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
              {fileChoosen ? `${dnaFiles.name}` : "Upload file"}
              <input
                id="dna"
                type="file"
                name="dnaSequence"
                onChange={(e) => onChange(e)}
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
            <select
              name="diseaseId"
              onChange={(e) => setDiseasesId(e.target.value)}
            >
              {diseases.map((disease) => (
                <option key={disease.id} value={disease.id}>
                  {disease.disease_name}
                </option>
              ))}
            </select>
          </div>
        </div>
      </form>
    </div>
  );
};

export default TesDNA;
