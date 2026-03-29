import { BrowserRouter, Route, Routes } from "react-router-dom";
import { AccountPage } from "./pages/account-page/account-page";
import { DashboardPage } from "./pages/dashboard-page/dashboard-page";
import { GalleryPage } from "./pages/gallery-page/gallery-page";
import { HabitsPage } from "./pages/habits-page/habits-page";
import { JournalPage } from "./pages/journal-page/journal-page";

export default function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<DashboardPage />} />
        <Route path="/journal" element={<JournalPage />} />
        <Route path="/habits" element={<HabitsPage />} />
        <Route path="/gallery" element={<GalleryPage />} />
        <Route path="/account" element={<AccountPage />} />
      </Routes>
    </BrowserRouter>
  );
}
