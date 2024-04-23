const StartStep = ({
  setStep,
}: {
  setStep: React.Dispatch<
    React.SetStateAction<"start" | "handleFile" | "getMetrics">
  >;
}) => (
  <div className="flex flex-col items-center justify-center">
    <p className="text-2xl">Welcome</p>
    <h1 className="text-6xl font-bold">Stori Transaction handler</h1>
    <div className="mt-6">
      <button
        className="bg-stori-green text-white px-6 py-3 mx-2 rounded-full"
        onClick={() => setStep("handleFile")}
      >
        Choose File
      </button>
      <button
        className="bg-stori-green text-white px-6 py-3 mx-2 rounded-full"
        onClick={() => setStep("getMetrics")}
      >
        Get Metrics
      </button>
    </div>
  </div>
);

export default StartStep;
