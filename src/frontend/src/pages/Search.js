import { useState, useEffect } from "react";
import axios from "axios";
import { Box } from "@mui/material"

const formatDate = (date) => {
  // change format from yyyy-mm-dd to dd MONTHNAME yyyy
  const monthNames = [
    "January",
    "February",
    "March",
    "April",
    "May",
    "June",
    "July",
    "August",
    "September",
    "October",
    "November",
    "December",
  ];
  const dateObj = new Date(date);
  const day = dateObj.getDate();
  const monthIndex = dateObj.getMonth();
  const year = dateObj.getFullYear();
  return `${day} ${monthNames[monthIndex]} ${year}`;
};

const Search = () => {
  const [search, setSearch] = useState("");
  const [results, setResults] = useState([]);

  const handleChange = (e) => {
    setSearch(e.target.value);
    console.log(search)
  };

  useEffect(() => {
    const fetchData = async () => {
      const res = await axios(
        `https://enigmatic-brook-59106.herokuapp.com/api/prediction?q=${search}`
      );
      if (res.data != null) {
        setResults(res.data);
      } else {
        setResults([]);
      }
    };
    fetchData();
  }, [search]);

    return (
      <div>
        <h1 className="subtitle"> Search</h1>
        <div className="search">
              <input
                className="searchInput"
                type="text"
                name="penyakit"
                placeholder="Search..."
                // onChange={(e) => handleChange(e)}
                onKeyDown={
                  (e) => {
                    if (e.key === "Enter") {
                      handleChange(e);
                    }
                  }
                }
              />
        </div>
        <table>
          {results.map((disease) => {
            return (
              <Box className="searchResult">
                <tr key={disease.id}>
                  {formatDate(disease.date)} - {disease.patient_name} - {disease.disease_name}
                  - {disease.result ? "True" : "False"}
                </tr>
              </Box>
            )
          })}
        </table>
      </div>
    );
}

export default Search;
