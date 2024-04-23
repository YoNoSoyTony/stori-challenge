// FileDropper.tsx
import React, { useCallback, useState } from 'react';
import { useDropzone } from 'react-dropzone';
import Papa from 'papaparse';

interface FileDropperProps {
 onDataParsed: (data: any[]) => void;
}

const FileDropper: React.FC<FileDropperProps> = ({ onDataParsed }) => {
 const [error, setError] = useState<string | null>(null);

 const onDrop = useCallback((acceptedFiles: File[]) => {
    const file = acceptedFiles[0]; // Select only the first file
    Papa.parse(file, {
      header: true,
      complete: (results) => {
        if (results.errors.length > 0) {
          setError('Error parsing CSV file.');
        } else {
          onDataParsed(results.data);
        }
      },
      error: (err: any): void => {
        setError('Error reading file.');
        console.log(err)
      },
    });
 }, [onDataParsed]);

 const { getRootProps, getInputProps, isDragActive } = useDropzone({ onDrop });

 return (
    <div className="flex flex-col items-center justify-center p-4 border-dashed border-2 border-gray-400 rounded-lg">
      <div {...getRootProps()} className="w-full h-32 flex flex-col items-center justify-center cursor-pointer">
        <input {...getInputProps()} />
        {isDragActive ? (
          <p>Drop the files here ...</p>
        ) : (
          <p>Drag 'n' drop a CSV file here, or click to select a file</p>
        )}
      </div>
      {error && <div className="text-red-500 mt-2">{error}</div>}
    </div>
 );
};

export default FileDropper;
