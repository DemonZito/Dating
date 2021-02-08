const version = "46cede8606bfe3e17d074ae8a651d07b";self.addEventListener('install', function(event) {
		self.skipWaiting();
		caches.delete("dynamic");
  event.waitUntil(
    caches.open("assets").then(function(cache) {
      return cache.addAll(
        ["/", "/assets/fixedsys.ttf"]
      );
    }).catch(function(e) {
		console.log("Couldn't install because: ", e);
	})
  );
});

self.addEventListener('fetch', event => event.respondWith(cacheThenNetwork(event)));

async function cacheThenNetwork(event) {
	let request = event.request;

	const assets = await caches.open("assets");

	//Try load a cached asset first.
	const CachedAsset = await assets.match(request);
	if (CachedAsset) return CachedAsset;

	//Get the request from the network.
	try {
		let clone = request.clone();
		clone.url = request.url+"?="+Math.random();
		const NetworkReponse = await fetch(clone, {cache: "no-store"});
		if (request.method == "GET" && NetworkReponse.status == 200) {
			const dynamic = await caches.open("dynamic");
			dynamic.put(request, NetworkReponse.clone());
		}
		return NetworkReponse;
	} catch (e) {
		//Try the dynamic cache.
		if (request.method == "GET") {
			const dynamic = await caches.open("dynamic");
			const CachedDynamic = await dynamic.match(request);
			if (CachedDynamic) return CachedDynamic;
		}

		return new Response("404 not found", {
			status: 404,
		})
	}
}
