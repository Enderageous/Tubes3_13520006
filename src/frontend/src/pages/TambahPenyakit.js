import { component, useState, useRef } from "react";
import axios from "axios";
import { Button } from "@mui/material";

function TambahPenyakit() {
  const [fileChoosen, setFileChoosen] = useState(false);
  const [dnaSequence, setDnaSequence] = useState();
  const [diseaseName, setDiseasesName] = useState("");

  const handleSubmit = (event) => {
    event.preventDefault();
    // axios
    //   .post("https://enigmatic-brook-59106.herokuapp.com/api/disease", {
    //     diseaseName,
    //     dnaSequence,
    //   })
    //   .then((res) => console.log("Posting data", res))
    //   .catch((error) => console.log(error));

    axios({
      url: "https://enigmatic-brook-59106.herokuapp.com/api/disease",
      data: {
        diseaseName,
        dnaSequence,
      },
      headers: { "Access-Control-Allow-Origin": "*" },
      method: "post",
    })
      .then((res) => console.log("Posting data", res))
      .catch((error) => console.log(error));
  };

  const onChange = (e) => {
    if (e.target.files[0]) {
      setFileChoosen(true);
      setDnaSequence(e.target.files[0].name);
    } else {
      setFileChoosen(false);
    }
  };

  return (
    <div>
      <h1 className="subtitle"> Tambah Penyakit</h1>
      <div className="tambahPenyakit">
        <div className="twocolumn">
          <p>Nama Penyakit:</p>
          <form>
            <input
              id="penyakit"
              className="twoInput"
              type="text"
              name="diseaseName"
              value={diseaseName}
              onChange={(e) => setDiseasesName(e.target.value)}
              placeholder="penyakit..."
            />
          </form>
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
            {fileChoosen ? `${dnaSequence}` : "Upload file"}
            <input
              id="dna"
              type="file"
              name="dnaSequence"
              value={dnaSequence}
              onChange={(e) => setDnaSequence(e.target.value)}
              hidden
            ></input>
          </Button>
        </div>
        <div className="onecolumn">
          <form onSubmit={handleSubmit}>
            <button className="submitButton" type="submit">
              Submit
            </button>
          </form>
        </div>
      </div>
    </div>
  );
}

export default TambahPenyakit;
