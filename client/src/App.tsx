import { BrowserRouter as Router, Routes, Route } from "react-router-dom"
import Home from "./pages/Home"
import Team from "./pages/Team"
import { Layout } from "@/components/layout"
import EmergencyContacts from "./pages/Emergency-Contacts"
import Checklist from "./pages/Checklist.tsx"

function App() {
  return (
    <Router>
      <Layout>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/team/:id" element={<Team />} />
          <Route path="/emergency-contacts/:id" element={<EmergencyContacts />} />
          <Route path="/checklist/:id" element={<Checklist />} />
          <Route path="*" element={<div>404 Not Found</div>} />
        </Routes>
      </Layout>
    </Router>
  )
}

export default App
