import { useState } from "react";
import stori from "./assets/stori.png";
import StartStep from "./pages/StartStep";
import GetMetricsStep from "./pages/GetMetricsStep";
import HandleFileStep from "./pages/HandleFileStep";
import { ToastContainer } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

const App = () => {
  const [step, setStep] = useState<"start" | "handleFile" | "getMetrics">(
    "start"
  );

  return (
    <div className="w-full h-screen bg-stori-teal text-white">
      <div className="h-14 w-full bg-stori-teal flex flex-row items-center justify-center shadow-md">
        <img src={stori} alt="stori logo" className="h-12" />
      </div>
      {step !== "start" && (
        <div
          onClick={() => setStep("start")}
          className="flex items-center justify-center rounded-full bg-stori-green text-white p-4 cursor-pointer mt-4 ml-4"
          style={{ width: "50px", height: "50px" }}
        >
          {"←"}
        </div>
      )}
      <div className="flex flex-col items-center justify-center h-[60vh]">
        {step === "start" && <StartStep setStep={setStep} />}
        {step === "handleFile" && <HandleFileStep />}
        {step === "getMetrics" && <GetMetricsStep />}
      </div>
      <div className="h-14 w-full bg-stori-teal flex flex-row items-center justify-center absolute bottom-0">
        <p className="text-white px-2">© 2024 Stori</p>
        <p className="text-white px-2">@yonosoytony</p>
      </div>
      <ToastContainer />
    </div>
  );
};

export default App;
