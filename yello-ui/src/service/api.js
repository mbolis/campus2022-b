export async function getAllLists() {
    const resp = await fetch("/lists");
    const result = await resp.json();
    return result.lists;
}

export async function addNewList(list) {
    const resp = await fetch("/lists", {
        method: "POST",
        headers: { "content-type": "application/json" },
        body: JSON.stringify(list),
    });
    return await resp.json();
}