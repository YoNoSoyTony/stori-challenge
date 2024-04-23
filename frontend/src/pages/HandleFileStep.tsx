import FileDropper from '../components/FileDropper'

const HandleFileStep = () => {
  return (
  <FileDropper onDataParsed={(data) => console.log(data)} />
  )
}
export default HandleFileStep;
