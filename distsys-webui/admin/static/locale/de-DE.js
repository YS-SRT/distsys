import {registerLocale as register} from 'amis-core';

amisRequire('amis').registerLocale('de-DE', {
  'Action.countDown': '${timeLeft} warten',
  'Alert.info': 'Systeminformationen',
  'asc': 'Aufsteigend',
  'App.home': 'Startseite',
  'App.navigation': 'Navigation',
  'Calendar.datepicker': 'Auswahl des Datums',
  'Calendar.yearmonth': 'YYYY MM',
  'Calendar.year': 'YYYY',
  'Calendar.begin': 'beginnt',
  'Calendar.end': 'ende',
  'Calendar.beginAndEnd': 'b/e',
  'Calendar.toast': 'Außerhalb des Datumsbereichs',
  'Calendar.startPick': 'Wählen Sie Startzeit',
  'Calendar.endPick': 'Wählen Sie Endzeit',
  'Schedule': 'Zeitplan',
  'Time': 'Zeit',
  'Content': 'Inhalt',
  'cancel': 'Abbrechen',
  'more': 'mehr',
  'Card.dragTip': 'Obere Schaltfläche zum Sortieren ziehen',
  'Card.toggleDrag': 'Zum Sortieren umschalten',
  'City.street': 'Straße eingeben',
  'clear': 'Löschen',
  'ColorPicker.placeholder': 'Farbe auswählen',
  'SubForm.add': 'Neu',
  'add': 'Neu',
  'Combo.dragDropSort': 'Zum Sortieren ziehen',
  'Combo.invalidData': 'Ungültige Daten, bitte entfernen',
  'Combo.maxLength':
    'Maximale Anzahl ist {{MaxLength}}}. Löschen Sie einige Elemente.',
  'Combo.minLength':
    'Mindestens {{minLength}} erfoderlich. Fügen Sie weitere hinzu.',
  'Combo.type': 'Typ',
  'confirm': 'Bestätigen',
  'Copyable.tip': 'Kopieren',
  'CRUD.exportCSV': 'In CSV exportieren',
  'CRUD.exportExcel': 'In Excel exportieren',
  'CRUD.exportExcelTemplate': 'In Excel-Vorlage exportieren',
  'CRUD.fetchFailed': 'Fehler beim Abrufen',
  'CRUD.filter': 'Filtern',
  'CRUD.selected': 'Ausgewählte {{total}} Elemente: ',
  'CRUD.invalidArray': '"data.items" muss ein Array sein',
  'CRUD.invalidData': '"data" ist leer',
  'CRUD.loadMore': 'Weitere laden',
  'CRUD.loadMoreDisableTip': 'Keine Daten oder letzte Seite',
  'CRUD.perPage': 'Pro Seite',
  'CRUD.stat': '{{page}} von {{lastPage}} insgesamt: {{total}}.',
  'CRUD.paginationGoText': 'Wechseln zu',
  'CRUD.paginationPageText': 'Seite',
  'PaginationWrapper.placeholder': 'Textkörper konfigurieren',
  'Pagination.select': '{{count}} items/page',
  'Pagination.goto': 'goto',
  'Pagination.go': 'GO',
  'Pagination.totalPage': 'total {{lastPage}} pages',
  'Pagination.totalCount': 'total {{total}} items',
  'Date.titleYear': '',
  'Date.titleMonth': '',
  'Date.titleQuarter': '',
  'Date.titleDate': '',
  'Date.titleTime': '',
  'Date.daysago': 'Vor {{days}} Tag(en)',
  'Date.dayslater': '{{days}} Tag(e) später',
  'Date.endOfMonth': 'Letzter Tag des Monats',
  'Date.endOfLastMonth': 'Letzer Tag des letzten Monats',
  'Date.endOfWeek': 'Samstag',
  'Date.hoursago': 'Vor {{hours}} Stunde(n)',
  'Date.hourslater': '{{hours}} Stunde(n) später',
  'Date.invalid': 'Ungültiges Datum',
  'Number.invalid': 'Ungültige Zahl',
  'Date.monday': 'Montag',
  'Date.monthsago': 'Vor {{months}} Monat(en)',
  'Date.monthslater': '{{months}} Monat(e) später',
  'Date.now': 'Jetzt',
  'Date.placeholder': 'Datum wählen',
  'Date.quartersago': 'Vor {{quarters}} Quartal(en)',
  'Date.quarterslater': '{{quarters}} Quartal(e) später',
  'Date.startOfLastMonth': 'Erster Tag des letzten Monats',
  'Date.startOfLastQuarter': 'Erster Tag des letzten Quartals',
  'Date.startOfMonth': 'Erster Tag des Monats',
  'Date.startOfQuarter': 'Erster Tag des Quartals',
  'Date.today': 'Heute',
  'Date.tomorrow': 'Morgen',
  'Date.weeksago': 'vor {{weeks}} Woche',
  'Date.weekslater': '{{weeks}} Wochen später',
  'Date.yesterday': 'Gestern',
  'dateformat.year': 'YYYY',
  'DateRange.daysago': 'letzten {{days}} Tage',
  'DateRange.dayslater': 'innerhalb von {{days}} Tagen',
  'DateRange.weeksago': 'letzten {{weeks}} Wochen',
  'DateRange.weekslater': 'innerhalb von {{weeks}} Wochen',
  'DateRange.monthsago': 'letzten {{months}} Monate',
  'DateRange.monthslater': 'innerhalb von {{months}} Monaten',
  'DateRange.quartersago': 'letzten {{quarters}} Quartale',
  'DateRange.quarterslater': 'innerhalb von {{quarters}} Quartalen',
  'DateRange.yearsago': 'letzten {{years}} Jahre',
  'DateRange.yearslater': '{{years}} Jahren',
  'DateRange.hoursago': 'letzten {{hours}} Stunden',
  'DateRange.hourslater': 'innerhalb von {{hours}} Stunden',
  'DateRange.1dayago': 'Vor 1 Tag',
  'DateRange.1daysago': 'Vor 1 Tag',
  'DateRange.7daysago': 'Vor 7 Tagen',
  'DateRange.30daysago': 'Vor 30 Tagen',
  'DateRange.90daysago': 'Vor 90 Tagen',
  'DateRange.lastMonth': 'Letzer Monat',
  'DateRange.lastWeek': 'Letzte Woche',
  'DateRange.lastQuarter': 'Letztes Quartal',
  'DateRange.placeholder': 'Datumsbereich wählen',
  'YearRange.placeholder': 'Datumsbereich wählen',
  'DateRange.thisWeek': 'Diese Woche',
  'DateRange.thisMonth': 'Diesen Monat',
  'DateRange.thisQuarter': 'Dieses Quartal',
  'DateRange.thisYear': 'Dieses Jahr',
  'DateRange.lastYear': 'letztes Jahr',
  'DateRange.valueConcat': ' bis ',
  'DateTime.placeholder': 'Datum auswählen',
  'delete': 'Löschen',
  'deleteConfirm': 'Möchten Sie wirklich löschen?',
  'deleteFailed': 'Fehler beim Löschen',
  'desc': 'Absteigend',
  'Dialog.close': 'Schließen',
  'Dialog.title': 'Titel',
  'Embed.invalidRoot': 'Ungültiger Root-Selektor',
  'Embed.downloading': 'Download starten',
  'fetchFailed': 'Fehler beim Abrufen der API',
  'File.continueAdd': 'Hinuzufügen fortsetzen',
  'File.dragDrop': 'Dateien per Drag & Drop hier ablegen',
  'File.clickUpload': 'Klicken Sie hier zum Hochladen',
  'File.helpText': 'Hilfedokumentation',
  'File.errorRetry':
    'Fehler beim Hochladen der Datei, bitte versuchen Sie es erneut.',
  'File.failed': 'Fehlerhafte Dateien',
  'File.invalidType': '{{files}} entspricht nicht Typ `{{accept}}`',
  'File.maxSize':
    '{{filename}} überschreitet die maximale Größe von {{maxSize}}',
  'File.pause': 'Hochladen anhalten',
  'File.repick': 'Erneut suswählen',
  'File.result':
    'Erfolgreich hochgeladene Dateien: {{uploaded}}, nicht hochgeladene Dateien: {{failed}}',
  'File.retry': 'Wiederholen',
  'File.sizeLimit': 'Die maximale Dateigröße ist {{maxSize}}',
  'File.start': 'Hochladen beginnen',
  'File.upload': 'Hochladen',
  'File.uploadFailed': 'Zurückgegebene Daten der Upload-API sind leer',
  'File.uploading': 'Wird hochgeladen...',
  'File.imageAfterCrop': 'Beschnittenes Bild',
  'FormItem.autoFillLoadFailed':
    'Die Schnittstelle hat einen Fehler zurückgegeben, bitte sorgfältig prüfen',
  'FormItem.autoFillSuggest': 'Referenzdateneingabe',
  'Form.loadOptionsFailed':
    'Optionen wurden auf folgendem Grund nicht geladen: {{reason}}',
  'Form.submit': 'Absenden',
  'Form.title': 'Formular',
  'Form.unique': 'Aktueller Wert ist nicht eindeutig',
  'Form.validateFailed': 'Fehler bei der Überprüfung der Formulareingabe',
  'Form.nestedError': 'Form kann nicht als Nachkomme von Form erscheinen',
  'Form.rules.message':
    'Die gemeinsame Überprüfung von Formularelementen ist fehlgeschlagen',
  'Iframe.invalid': 'Ungültige Iframe-URL',
  'Iframe.invalidProtocol':
    'HTTP-URL-Iframe kann nicht in https verwendet werden',
  'Image.dragTip': 'Zum Sortieren ziehen',
  'Image.upload': 'Bild hochladen',
  'Image.configError':
    'Es können nur eine Beschneidung oder mehrere festgelegt werden',
  'Image.crop': 'Bild beschneiden',
  'Image.dragDrop': 'Bilder per Drag & Drop hier ablegen',
  'Image.height': 'Höhe: {{height}} Pixel',
  'Image.limitMax': 'Minimale Bildgröße ist {{info}}',
  'Image.limitMin': 'Maximale Bildgröße ist {{info}}',
  'Image.limitRatio':
    'Laden Sie das Bild mit dem Seitenverhältnis {{ration}} hoch',
  'Image.pasteTip': 'Sie können das Bild aus der Zwischenablage einfügen',
  'Image.placeholder':
    'Klicken Sie, um das Bild einzufügen, oder ziehen Sie es in diesen Bereich.',
  'Image.size': 'size: ({{width}} Pixel x {{height}} Pixel)',
  'Image.sizeNotEqual':
    'Das ausgwählte Bild entspricht nicht den Größenanforderungen {{info}}',
  'Image.width': 'Weite: {{width}} Pixel',
  'Image.zoomIn': 'Vergrößern',
  'Log.mustHaveSource': 'Quelle muss in der Konfiguration vorhanden sein',
  'Log.showLineNumber': 'Zeilennummer anzeigen',
  'Log.notShowLineNumber': 'Zeilennummer ausblenden',
  'Log.expand': 'Entfalten',
  'Log.collapse': 'Falten',
  'link': 'Link',
  'loading': 'Wird geladen...',
  'loadingFailed': 'Das Laden ist fehlgeschlagen',
  'LocationPicker.placeholder': 'Wählen Sie einen Ort',
  'LocationPicker.getLocation':
    'Klicken Sie hier, um Standortinformationen zu erhalten',
  'Month.placeholder': 'Wählen Sie einen Monat',
  'Nav.sourceError': 'Fehler beim Abrufen des Links',
  'networkError':
    'Fehler beim Netzwerkzugriff oder fehlende CORS-Konfiguration',
  'noResult': 'Keine Ergebnisse',
  'NumberInput.placeholder': 'Geben Sie eine Zahl ein',
  'Options.addPlaceholder': 'Geben Sie einen Namen ein',
  'Options.deleteAPI': '"deleteAPI" erforderlich',
  'Options.editLabel': 'Bearbeiten {{label}}',
  'Options.label': 'Option',
  'Options.createFailed': 'Erstellen fehlgeschlagen',
  'Options.retry':
    "Laden fehlgeschlagen '{{reason}}', klicken Sie auf Wiederholen",
  'placeholder.empty': '<Empty>',
  'placeholder.enter': 'Eingabe',
  'placeholder.noData': 'Keine Daten',
  'placeholder.noOption': 'Keine Option',
  'placeholder.selectData': 'Daten auswählen',
  'Quarter.placeholder': 'Quartal auswählen',
  'Repeat.pre': 'Pro',
  'reset': 'Zurücksetzen',
  'save': 'Konservierung',
  'saveFailed': 'Fehler beim Speichern',
  'saveSuccess': 'Erfolgreich gespeichert',
  'search': 'Suchen',
  'searchHistory': 'Suchverlauf',
  'searchResult': 'Suchergebnis',
  'Checkboxes.selectAll': 'Alle auswählen/abwählen',
  'Select.checkAll': 'Alle markieren',
  'Select.clear': 'Löschen',
  'Select.upload': 'Wieder hochladen',
  'Select.clearAll': 'Alle löschen',
  'Select.createLabel': 'Neue Option',
  'Select.placeholder': 'Auswählen',
  'Select.searchPromptText': 'Eingeben zum Suchen',
  'Select.selected': 'Ausgewählt',
  'sort': 'Sortieren',
  'SubForm.button': 'Configurieren',
  'SubForm.editDetail': 'Details bearbeiten',
  'System.error': 'Systemfehler',
  'System.notify': 'Systembenachrichtigung',
  'System.copy': 'Inhalt kopiert',
  'System.requestError': 'Anfragefehler: ',
  'System.requestErrorStatus': 'Anfragefehler, Statuscode:',
  'Table.addRow': 'Zeile hinzufügen',
  'Table.copyRow': 'Zeile kopieren',
  'Table.columnsVisibility':
    'Klicken, um die Sichtbarkeit der Spalten zu steuern',
  'Table.deleteRow': 'Aktuele Zeile löschen',
  'Table.discard': 'Verwerfen',
  'Table.dragTip': 'Schaltfläche links zum Sortieren ziehen',
  'Table.editing': 'Sie müssen die Bearbeitung beenden.',
  'Table.editRow': 'Aktuelle Zeile bearbeiten',
  'Table.modified': 'Es wurden {{modified}} Datensätze geändert, Sie können:',
  'Table.moved':
    'Bei {{moved}} Datensätzen wurde die Reihenfolge geändert, Sie können:',
  'Table.operation': 'Vorgang',
  'Table.playload': 'Nutzlast muss vorhanden sein',
  'Table.startSort': 'Klicken, um Sortierung zu starten',
  'Table.valueField': 'valueField muss vorhanden sein',
  'Table.index': 'Index',
  'Table.add': 'Neu',
  'Table.subAddRow': 'Unterzeile hinzufügen',
  'Table.addButtonDisabledTip':
    'Reichen Sie bei der Inhaltsbearbeitung zuerst ein und erstellen Sie dann eine neue Option',
  'Table.toggleColumn': 'Spalten anzeigen',
  'Table.searchFields': 'Abfragefelder setzen',
  'Tag.placeholder': 'Noch kein Tag',
  'Tag.tip': 'Kürzlich verwendetes Tag',
  'Text.add': 'Neu {{label}}',
  'Time.placeholder': 'Zeit auswählen',
  'Transfer.configError': 'Konfigurationsfehler',
  'Transfer.refreshIcon': 'Zum Aktualisieren klicken',
  'Transfer.searchKeyword': 'Stichwörter eingeben',
  'Transfer.available': 'Verfügbar',
  'Transfer.selectd': 'Ausgewählt',
  'Transfer.selectFromLeft': 'Von links auswählen',
  'Tree.addChild': 'Untergeordnetes Element hinzufügen',
  'Tree.addRoot': 'Stammknoten hinzufügen',
  'Tree.editNode': 'Diesen Knoten bearbeiten',
  'Tree.invalidArray':
    'Daten.Optionen, Daten.Elemente oder Daten müssen ein Array sein',
  'Tree.removeNode': 'Diesen Knoten entfernen',
  'Tree.root': 'Stamm',
  'validate.equals': 'Wert muss $1 sein',
  'validate.equalsField': 'Wert muss $1 sein',
  'validate.gt': 'Geben Sie einen Wert ein, der größer ist als $1',
  'validate.isAlpha': 'Geben Sie Buchstaben ein',
  'validate.isAlphanumeric': 'Geben Sie Buchstaben oder Zahlen ein.',
  'validate.isEmail': 'E-Mail-Format ist falsch',
  'validate.isFloat': 'Geben Sie einen Gleitkommawert ein',
  'validate.isId': 'Ungültige ID-Kartennummer',
  'validate.isInt': 'Geben Sie eine ganze Zahl ein',
  'validate.isJson': 'Ungültiges JSON-Format',
  'validate.isLength':
    'Vergewissern Sie sich, dass die Länge des Inhalts $1 ist',
  'validate.isNumeric': 'Geben Sie eine Nummer ein',
  'validate.isPhoneNumber': 'Ungültige Telefonnummer',
  'validate.isRequired': 'Dies ist erforderlich',
  'validate.isTelNumber': 'Ungültige Telefonnummer',
  'validate.isUrl': 'Falsches URL-Format',
  'validate.isUrlPath':
    'Sie können nur Buchstaben, Zahlen, "-" und "_" eingeben.',
  'validate.isWords': 'Geben Sie ein Wort ein',
  'validate.isZipcode': 'Ungültige Postleitzahl',
  'validate.lt': 'Geben Sie einen Wert ein, der kleiner ist als $1',
  'validate.matchRegexp':
    'Das Format ist nicht korrekt. Geben Sie den Inhalt mit der Regel `${1| raw}` ein.',
  'validate.maximum':
    'Der Eingabewert überschreitet den maximalen Wert von $1.',
  'validate.maxLength':
    'Kontrollieren Sie die Länge des Inhalts. Geben Sie nicht mehr als $1 Buchstaben ein.',
  'validate.minimum': 'Der Eingabewert ist kleiner als der Mindestwert von $1.',
  'validate.minLength': 'Geben Sie weitere Zeichen ein, mindestens $1.',
  'validate.array.minLength':
    'Bitte fügen Sie weitere Mitglieder hinzu, mindestens $1 Mitglieder',
  'validate.array.maxLength':
    'Bitte kontrollieren Sie die Anzahl der Mitglieder, die $1 nicht überschreiten darf',
  'validate.notEmptyString': 'Geben Sie nicht nur Leerzeichen ein.',
  'validate.isDateTimeSame':
    'Der aktuelle Datumswert ist ungültig, bitte geben Sie denselben Datumswert wie $1 ein',
  'validate.isDateTimeBefore':
    'Der aktuelle Datumswert ist ungültig, bitte geben Sie einen Datumswert vor $1 ein',
  'validate.isDateTimeAfter':
    'Der aktuelle Datumswert ist ungültig, bitte geben Sie nach $1 einen Datumswert ein',
  'validate.isDateTimeSameOrBefore':
    'Der aktuelle Datumswert ist ungültig. Bitte geben Sie einen Datumswert ein, der gleich oder älter als $1 ist',
  'validate.isDateTimeSameOrAfter':
    'Der aktuelle Datumswert ist ungültig. Bitte geben Sie einen Datumswert ein, der gleich oder nach $1 ist',
  'validate.isDateTimeBetween':
    'Der aktuelle Datumswert ist ungültig, bitte geben Sie einen Datumswert zwischen $1 und $2 ein',
  'validate.isTimeSame':
    'Der aktuelle Zeitwert ist ungültig, bitte geben Sie denselben Zeitwert wie 1 $ ein',
  'validate.isTimeBefore':
    'Der aktuelle Zeitwert ist ungültig, bitte geben Sie einen Zeitwert vor $1 ein',
  'validate.isTimeAfter':
    'Der aktuelle Zeitwert ist ungültig, bitte geben Sie nach $1 einen Zeitwert ein',
  'validate.isTimeSameOrBefore':
    'Der aktuelle Zeitwert ist ungültig. Bitte geben Sie einen Zeitwert ein, der gleich oder älter als $1 ist',
  'validate.isTimeSameOrAfter':
    'Der aktuelle Zeitwert ist ungültig. Bitte geben Sie einen Zeitwert ein, der gleich oder nach $1 ist',
  'validate.isTimeBetween':
    'Der aktuelle Zeitwert ist ungültig, bitte geben Sie einen Zeitwert zwischen $1 und $2 ein',
  'validate.isVariableName':
    'Bitte geben Sie einen gültigen Variablennamen ein',
  'validateFailed': 'Fehler bei der Überprüfung',
  'Wizard.configError': 'Konfigurationsfehler',
  'Wizard.finish': 'Ende',
  'Wizard.next': 'Weiter',
  'Wizard.prev': 'Zurück',
  'Wizard.saveAndNext': 'Speichern & Weiter',
  'year-to-year': '{{from}} - {{to}}',
  'Year.placeholder': 'Wählen Sie ein Jahr',
  'reload': 'Neu laden',
  'rotate': 'Drehen',
  'rotate.left': 'Nach links drehen',
  'rotate.right': 'Drehe nach rechts',
  'zoomIn': 'Vergrößern',
  'zoomOut': 'Verkleinern',
  'scale.origin': 'Originalmaße',
  'Editor.fullscreen': 'Schirmfüllend Modus',
  'Editor.exitFullscreen': 'Zurücktreten Schirmfüllend Modus',
  'Condition.not': 'nicht',
  'Condition.and': 'und',
  'Condition.or': 'oder',
  'Condition.collapse': 'entfalten',
  'Condition.add_cond': 'und Bedingung',
  'Condition.add_cond_group': 'Bedingungsgruppe hinzufügen',
  'Condition.delete_cond_group': 'Konditionsgruppe löschen',
  'Condition.equal': 'gleich',
  'Condition.not_equal': 'ungleich',
  'Condition.less': 'weniger',
  'Condition.less_or_equal': 'weniger oder gleich',
  'Condition.greater': 'greater',
  'Condition.greater_or_equal': 'größder oder gleich',
  'Condition.between': 'zwischen',
  'Condition.not_between': 'nicht zwischen',
  'Condition.is_empty': 'leer',
  'Condition.is_not_empty': 'nicht leer',
  'Condition.like': 'beinhaltet',
  'Condition.not_like': 'beinhaltet nicht',
  'Condition.starts_with': 'beginnt mit',
  'Condition.ends_with': 'endet mit',
  'Condition.select_equals': 'gleich',
  'Condition.select_not_equals': 'nicht gleich',
  'Condition.select_any_in': 'beinhaltet',
  'Condition.select_not_any_in': 'beinhaltet nicht',
  'Condition.placeholder': 'Text einfügen',
  'Condition.cond_placeholder': 'Bedingung auswählen',
  'Condition.field_placeholder': 'Feld auswählen',
  'Condition.blank': 'leer',
  'Condition.expression': 'Ausdruck',
  'Condition.formula_placeholder': 'Bitte geben Sie eine Formel ein',
  'Condition.fun_error': 'Funktion ist undefiniert',
  'Condition.configured': 'Konfiguriert',
  'Condition.isRequired': 'Bedingung kann nicht leer sein',
  'InputTable.uniqueError': 'Column `{{label}}` unique validate failed',
  'Timeline.collapseText': 'Falten',
  'Timeline.expandText': 'Entfalten',
  'collapse': 'Falten',
  'expand': 'Entfalten',
  'FormulaEditor.btnLabel': 'Formel Bearbeiten',
  'FormulaEditor.title': 'Formel Editor',
  'FormulaEditor.run': 'Laufen',
  'FormulaEditor.sourceMode': 'Source Mode',
  'FormulaEditor.runContext': 'Run Context',
  'FormulaEditor.runResult': 'Run Result',
  'FormulaEditor.toggleAll': 'Expand All',
  'FormulaEditor.variable': 'Variable',
  'FormulaEditor.function': 'Funktion',
  'FormulaEditor.invalidData':
    'Überprüfungsfehler, position or reason is {{err}}',
  'FormulaEditor.invalidValue':
    'Überprüfungsfehler, reason is Falsches Werteformat',
  'pullRefresh.pullingText': 'Zum Aktualisieren nach unten ziehen...',
  'pullRefresh.loosingText': 'Zum Aktualisieren freigeben...',
  'pullRefresh.loadingText': 'Laden...',
  'pullRefresh.successText': 'Laden erfolgreich',
  'Picker.placeholder': 'Klicken Sie rechts auf das Symbol',
  'UserSelect.edit': 'bearbeiten',
  'UserSelect.save': 'Konservierung',
  'UserSelect.resultSort': 'Ergebnissortierung auswählen',
  'UserSelect.selected': 'Ausgewählt',
  'UserSelect.clear': 'leer',
  'UserSelect.sure': 'Submit',
  'SchemaType.string': 'String',
  'SchemaType.number': 'Zahl',
  'SchemaType.integer': 'Ganzzahl',
  'SchemaType.object': 'Objekt',
  'SchemaType.array': 'Array',
  'SchemaType.boolean': 'Boolesch',
  'SchemaType.null': 'Null',
  'SchemaType.any': 'Any',
  'JSONSchema.type': 'Typ',
  'JSONSchema.required': 'Erforderlich',
  'JSONSchema.title': 'Titel',
  'JSONSchema.default': 'Standard',
  'JSONSchema.description': 'Beschreibung',
  'JSONSchema.key': 'Schlüssel',
  'JSONSchema.array_items': 'Artikel',
  'JSONSchema.members': 'Mitglieder',
  'JSONSchema.key_duplicated': 'Schlüssel existiert bereits',
  'TimeNow': 'Jetzt',
  'Steps.step': 'Schritt {{index}}',
  'FormulaInput.True': 'Treu',
  'FormulaInput.False': 'Falsch',
  'Signature.clear': 'leer',
  'Signature.undo': 'widerrufen',
  'Signature.confirm': 'bestätigen',
  'Signature.cancel': 'Abbrechen',
  'Signature.embedLabel': 'Klicken Sie zum Signieren'
});

