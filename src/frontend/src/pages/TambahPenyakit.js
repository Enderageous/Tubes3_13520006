import { Component, useState } from "react";
import { axios } from "axios";

class TambahPenyakit extends Component {
  state = {
    name: "",
  };

  setnama(event) {
    this.setState({ name: event.target.value });
  }

  handleSubmit = (event) => {
    event.preventDefault();
    const name = { name: this.state.name };

    axios
      .post(`https://jsonplaceholder.typicode.com/users`, { name })
      .then((res) => {
        console.log(res);
        console.log(res.data);
      });
  };

  onChange(e) {
    let files = e.target.files;
    let reader = new FileReader();
    reader.readAsDataURL(files[0]);

    reader.onload = (e) => {
      console.warn("data uploaded ", e.target.result); //checking only, will be deleted
    };
  }
  render() {
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
                name="name"
                onChange={this.setnama}
                placeholder="nama..."
              />
            </form>
          </div>
          {/* <div className="twocolumn">
            <p>Sequence DNA:</p>
            <input
              className="uploadButtonPenyakit"
              type="file"
              name="sequence_dna"
              onChange={(e) => this.onChange(e)}
            ></input>
          </div> */}
          <div className="onecolumn">
            <form onSubmit={this.handleSubmit}>
              <button className="submitButton" type="submit">
                Submit
              </button>
            </form>
          </div>
        </div>
      </div>
    );
  }
}

export default TambahPenyakit;
