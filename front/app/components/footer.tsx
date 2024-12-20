export default function Footer() {
  return (
    <footer className="bg-gray-800 text-white py-4 fixed bottom-0 w-full">
      <div className="container mx-auto text-center">
        <p>
          &copy; {new Date().getFullYear()} My Blog App. All rights reserved.
        </p>
      </div>
    </footer>
  );
}
