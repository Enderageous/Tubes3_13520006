import React from "react";
import axios from "axios";

class Search extends React.Component {
  state = {
    diseases: [],
    dates: [],
    names: [],
    results: [],
  };

  componentDidMount() {
    axios
      .get(`https://enigmatic-brook-59106.herokuapp.com/api/prediction`)
      .then((res) => {
        const diseases = res.data;
        const dates = res.data;
        const names = res.data;
        const results = res.data;
        this.setState({ diseases, dates, names, results });
      });
  }

  setSearch(event) {
    this.setState({ value: event.target.value });
  }

  render() {
    return (
      <div>
        <h1 className="subtitle"> Search</h1>
        <div className="search">
          <form>
            {
              <input
                className="searchInput"
                type="text"
                name="penyakit"
                placeholder="Search..."
                value={this.state.value}
                onChange={this.setSearch}
              />
            }
          </form>
        </div>
        <table className="searchResult">
          {this.state.diseases.map((disease) => (
            <tr key={disease.id}>
              {disease.date} - {disease.patient_name} - {disease.disease_name} -{" "}
              {disease.result ? "True" : "False"}
            </tr>
          ))}
        </table>
      </div>
    );
  }
}

export default Search;
