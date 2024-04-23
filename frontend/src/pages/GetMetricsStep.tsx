import React, { useState } from "react";
import axios from "axios";
import { TextField } from "@mui/material";
import { toast } from "react-toastify";

const GetMetricsStep = () => {
  const [email, setEmail] = useState("");

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();
    try {
      const response = await axios.post(
        "https://cld7djirp5.execute-api.us-east-1.amazonaws.com/default/sendEmail",
        {
          email,
        },
        {
          headers: {
            "Content-Type": "application/json",
          },
        }
      );
      toast.success(JSON.stringify(response.data));
    } catch (error) {
      toast.error("An error occurred while getting your account metrics.");
    }
  };

  return (
    <div className="flex flex-col items-center justify-center bg-white text-black rounded-xl p-6 shadow-lg w-full md:w-1/3">
      <div className="w-full pb-2 text-xl text-bold">
        <h2>Lets get your metrics</h2>
      </div>
      <TextField
        required
        label="Email"
        variant="outlined"
        value={email}
        onChange={(e) => setEmail(e.target.value)}
        className="w-full py-2 m-4"
      />
      <div className="h-2" />
      <button
        className="bg-stori-green text-white px-6 py-3 mx-2 my-0 rounded-full"
        onClick={handleSubmit}
      >
        Get Metrics
      </button>
    </div>
  );
};

export default GetMetricsStep;
