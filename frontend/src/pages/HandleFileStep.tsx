import React, { useState, useCallback, useEffect } from "react";
import { useDropzone } from "react-dropzone";
import Papa from "papaparse";
import { toast } from "react-toastify";
import TextField from "@mui/material/TextField";
import axios from 'axios';

interface Transaction {
  Id: string;
  Date: string;
  Transaction: string;
}
const submitTransactions = (transactions: any[]) => {
  axios.post('https://cld7djirp5.execute-api.us-east-1.amazonaws.com/default/handleCSV', {transactions: transactions})
     .then(response => {
       console.log(response);
       transactions.length !== 0 && toast.success("Data processed successfully.");
     })
     .catch(error => {
       console.error(error);
       console.log(transactions);
       transactions.length !== 0 && toast.error("Error processing data.");
     });
 };
 
const HandleFileStep: React.FC = () => {
  const [email, setEmail] = useState<string>("");
  const [selectedFile, setSelectedFile] = useState<File | null>(null);

  const onDrop = useCallback((acceptedFiles: File[]) => {
    const file = acceptedFiles[0];
    setSelectedFile(file);
  }, []);

  const { getRootProps, getInputProps, isDragActive } = useDropzone({ onDrop });

  const handleEmailChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setEmail(event.target.value);
  };

  const convertDateToMonthName = (transactions: Transaction[]): Transaction[] => {
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

    return transactions.map((transaction) => {
      const [monthNumber, restOfDate] = transaction.Date.split("/");
      const monthName = monthNames[parseInt(monthNumber) - 1];
      transaction.Date = `${monthName}/${restOfDate}`;
      return transaction;
    });
  };

  const handleProcessData = () => {
    if (!selectedFile) {
      toast.error("No file selected.");
      return;
    }

    Papa.parse(selectedFile, {
      header: true,
      dynamicTyping: true,
      complete: (results) => {
        if (results.errors.length > 0) {
          console.log(results.errors);
          toast.error("Error parsing CSV file.");
        } else {
          const processedTransactions = results.data.map((transaction: any) => ({
            email: email, 
            amount: parseFloat(transaction.Transaction), 
            month: convertDateToMonthName([transaction])[0].Date.split('/')[0] 
          }));
          submitTransactions(processedTransactions);
        }
      },
      error: (err: any): void => {
        toast.error("Error reading file.");
        console.log(err);
      },
    });
  };

  const handleClearFile = () => {
    setSelectedFile(null);
  };

  const fileName = selectedFile ? selectedFile.name : "";

  useEffect(() => {
    submitTransactions([]);
  }
  ,[]);

  return (
    <div className="flex flex-col items-center justify-center bg-white text-black rounded-xl p-6 shadow-lg w-full md:w-1/3">
      <div className="flex flex-col items-center justify-center">
        <div className="flex flex-col items-center justify-center p-4 border-dashed border-2 border-gray-400 rounded-lg">
          <div
            {...getRootProps()}
            className="w-full h-32 flex flex-col items-center justify-center cursor-pointer"
          >
            <input {...getInputProps()} />
            {isDragActive ? (
              <p>Drop the files here ...</p>
            ) : (
              <p>Drag 'n' drop a CSV file here, or click to select a file</p>
            )}
            <div className="w-full flex flex-row justify-between items-center border rounded-full pl-5 m-3">
              <p>{fileName}</p>
              {selectedFile && (
                <button
                  onClick={handleClearFile}
                  className="flex items-center justify-center rounded-full bg-red-500 text-white p-4 cursor-pointer"
                  style={{ width: "50px", height: "50px" }}
                >
                  {"✕"}
                </button>
              )}
            </div>
          </div>
        </div>
        <div className="h-2"/>
        <div className="w-full">

        </div>
        <TextField
          label="Email"
          variant="standard"
          fullWidth
          value={email}
          onChange={handleEmailChange}
        />
        <div className="h-2"/>
        <button
          className="bg-stori-green text-white px-6 py-3 mx-2 my-0 rounded-full"
          onClick={handleProcessData}
        >
          Process Data
        </button>
      </div>
    </div>
  );
};

export default HandleFileStep;
