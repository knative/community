This is a tool for migrating all the documents from the
[current Google Drive](https://drive.google.com/drive/folders/0APnJ_hRs30R2Uk9PVA)
which is owned by Google.com to a
[new Google Drive](https://drive.google.com/drive/folders/0AM-QGZJ-HUA8Uk9PVA)
which is owned by an independent GSuite organization (`knative.team`).

Moving to a GSuite domain which is independent of the Google.com domain will
allow us more flexibility in sharing settings, including the option to "publish
to the web", as well as better ability for non-Googlers to perform
administrative actions like sharing calendars.

This particular tool should do three things:

- [ ] [Enumerate all the documents](https://developers.google.com/drive/api/v3/reference/files/list)
      in a source GSuite shared drive
- [ ] [Make a copy](https://developers.google.com/drive/api/v3/reference/files/copy)
      of all the documents
- [ ] [Update all docs with a header](https://developers.google.com/docs/api/how-tos/move-text#inserting_text)
      indicating the new preferred location of the document.
