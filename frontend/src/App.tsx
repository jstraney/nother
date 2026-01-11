import { EditorLayout } from "./layouts/EditorLayout";
import { Provider } from "react-redux";
import { store } from "./stores/store";

function App() {
  return (
    <Provider store={store}>
      <div id="div__app">
        <EditorLayout />
      </div>
    </Provider>
  );
}

export default App;
