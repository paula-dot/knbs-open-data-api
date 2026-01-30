import { useEffect } from "react";
import { useCountyStore } from "./store/useCountyStore.js";

function App() {
    const { counties, isLoading, error, fetchCounties } = useCountyStore();

    useEffect(() => {
        // fetch once on mount
        fetchCounties();
    }, [fetchCounties]); // include fetchCounties to satisfy hooks rules

    if (isLoading) return <div className="p-10">Loading Kenya's data...</div>;
    if (error) return <div className="p-10 text-red-500">{error}</div>;

    return (
        <div className="max-w-4xl mx-auto p-10">
            <h1 className="text-3xl font-bold mb-10">Kenya Counties</h1>
            {counties.length === 0 ? (
                <div className="text-gray-500">No counties available yet.</div>
            ) : (
                <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                    {counties.map((county) => (
                        <div key={county.id} className="border p-4 rounded shadow hover:bg-gray-50">
                            <h2 className="text-lg font-bold mb-2">{county.code} - {county.name}</h2>
                            <p className="text-gray-600 text-sm">Province: {county.former_province ?? 'N/A'}</p>
                        </div>
                    ))}
                </div>
            )}
        </div>
    );
}

export default App;
