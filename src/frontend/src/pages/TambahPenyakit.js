import { useState } from "react";
import axios from "axios";
import { Button } from "@mui/material";

function TambahPenyakit() {
  const [fileChoosen, setFileChoosen] = useState(false);
  const [dnaSequence, setDnaSequence] = useState(null);
  const [diseaseName, setDiseasesName] = useState("");

  const handleSubmit = (event) => {
    event.preventDefault();
    const formData = new FormData();
    formData.append("disease_name", diseaseName);
    formData.append("dna_file", dnaSequence);
    // axios
    //   .post("https://enigmatic-brook-59106.herokuapp.com/api/disease", {
    //     diseaseName,
    //     dnaSequence,
    //   })
    //   .then((res) => console.log("Posting data", res))
    //   .catch((error) => console.log(error));

    axios({
      url: "https://enigmatic-brook-59106.herokuapp.com/api/disease",
      data: formData,
      headers: { "Content-Type": "multipart/form-data" },
      method: "post",
    })
      .then((res) => alert("Data penyakit berhasil ditambahkan"))
      .catch((error) => console.log(error));
  };

  const onChange = (e) => {
    if (e.target.files[0]) {
      setFileChoosen(true);
      setDnaSequence(e.target.files[0]);
    } else {
      setFileChoosen(false);
    }
  };

  return (
    <div>
      <form onSubmit={handleSubmit}>
      <h1 className="subtitle"> Tambah Penyakit</h1>
      <div className="tambahPenyakit">
        <div className="twocolumn">
          <p>Nama Penyakit:</p>
          <input
            id="penyakit"
            className="twoInput"
            type="text"
            name="diseaseName"
            onChange={(e) => setDiseasesName(e.target.value)}
            placeholder="penyakit..."
          />
        </div>
        <div className="twocolumn">
          <p>Sequence DNA:</p>
          <Button
            style={{
              borderRadius: 5,
              backgroundColor: "white",
              fontSize: "20px",
              color: "grey",
            }}
            variant="contained"
            component="label"
          >
            {fileChoosen ? `${dnaSequence.name}` : "Upload file"}
            <input
              id="dna"
              type="file"
              name="dnaSequence"
              onChange={(e) => onChange(e)}
              hidden
            ></input>
          </Button>
        </div>
        <div className="onecolumn">
          <button className="submitButton" type="submit">
            Submit
          </button>
        </div>
      </div>
      </form>
    </div>
  );
}

export default TambahPenyakit;
