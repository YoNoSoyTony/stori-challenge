import React, { useState } from "react";
import axios from "axios";
import { TextField } from "@mui/material";
import { toast } from "react-toastify";

const GetMetricsStep = () => {
  const [email, setEmail] = useState("");
  const [response, setResponse] = useState({} as Record<string, any>);

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
      if(response.data.transactions.length === 0) {
        toast.error("No transactions found for this email.");
        return;
      }
      else {
        setResponse(response.data);
      }
    } catch (error) {
      toast.error("An error occurred while getting your account metrics.");
    }
  };

  return (
    <div className="flex flex-col items-center justify-center bg-white text-black rounded-xl p-6 shadow-lg w-full md:w-1/3">
      {Object.keys(response).length === 0 && (
        <>
          <div className="w-full pb-2 text-xl text-bold">
            <h2>Lets get your metrics</h2>
          </div>
          <TextField
            required
            label="Email"
            variant="standard"
            value={email}
            onChange={(e) => setEmail(e.target.value.trim())}
            className="w-full py-2 m-4"
          />
          <div className="h-2" />
          <button
            className="bg-stori-green text-white px-6 py-3 mx-2 my-0 rounded-full"
            onClick={handleSubmit}
          >
            Get Metrics
          </button>
        </>
      )}
      {Object.keys(response).length > 0 && (
        <div className="w-full ">
          <h2 className="text-xl font-bold">Metrics</h2>
          <div className="flex flex-col items-start justify-start">
            <table className="w-full">
              {" "}
              {/* Apply w-full to make the table full width */}
              <thead>
                <tr>
                  <th>Metric</th>
                  <th>Value</th>
                </tr>
              </thead>
              <tbody>
                {response.balance !== undefined && (
                  <tr className="border-b border-gray-300">
                    <td className="text-base">Balance</td>
                    <td className="text-right">{response.balance}</td>
                  </tr>
                )}
                {response.positiveAverage !== undefined && (
                  <tr className="border-b border-gray-300">
                    <td className="text-base">Positive Average</td>
                    <td className="text-right">{response.positiveAverage}</td>
                  </tr>
                )}
                {response.negativeAverage !== undefined && (
                  <tr className="border-b border-gray-300">
                    <td className="text-base">Negative Average</td>
                    <td className="text-right">{response.negativeAverage}</td>
                  </tr>
                )}
              </tbody>
            </table>
          </div>
        </div>
      )}
      {response.transactionsByMonth && (
        <div className="flex flex-col items-start justify-start mt-4 mb-4 w-full">
          <h2 className="text-xl font-bold">Monthly Statements</h2>
          <table className="w-full">
            {" "}
            {/* Apply w-full to make the table full width */}
            <thead>
              <tr>
                <th>Month</th>
                <th>Transactions</th>
              </tr>
            </thead>
            <tbody>
              {Object.entries(response.transactionsByMonth).map(
                ([month, transactions]) => (
                  <tr key={month} className="border-b border-gray-300">
                    <td className="text-base">{month}</td>
                    <td className="text-right">{String(transactions)}</td>
                  </tr>
                )
              )}
            </tbody>
          </table>
        </div>
      )}
      {Object.keys(response).length > 0 && (
        <button
          className="bg-stori-green text-white px-6 py-3 mx-2 rounded-full"
          onClick={() => {
            setResponse({});
            setEmail("");
          }}
        >
          ← Get more Metrics
        </button>
      )}
    </div>
  );
};

export default GetMetricsStep;
