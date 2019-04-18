async patchCampus(input) { 
		
	if (input && input.PINProjectID) {
		
		try { 
		
			if (input.isSubmit) {
		
			input.SubmittedBy = currentUser.PWUserID
		
			}
		  
			let fields = input.UpdateField.split(',');
			let values = input.UpdateFieldValue.split(',');
			const currentUser = util.getLoggedInUser();
			let arrayUpdateFields = [];
			fields.forEach((field, index) => {
				const f = `${field}`;
				const v = `${values[index]}`;
				const r = f+v;
				arrayUpdateFields.push(r);
			});
			let command = [];
		
			command.push({key: APIVariable.CampusID, value: input.CampusID});
		
			command.push({key: APIVariable.strUpdateColumns, value: arrayUpdateFields.join()});
			command.push({key: APIVariable.ModifyBy, value: currentUser.PWUserID});
			
			const response = await dataAPI.getDataAPIProjectDB(QueryID.usp_AM_ProfileCampus_u_case1S_1, input.PINProjectID, command);
			if(response && response.result[0].affectedRows != 1) {
				
				return Promise.resolve(response);
			}
			if (input.isSubmit) {
		
			await this.updateReportSubmitAuditFields(input.PINProjectID, input);
		
			}
		
			return Promise.resolve({
				result: [{
				  error: false,CampusID: response.result
				}]
			});
		
		} catch (error) {
			logger.log('error',`ProfileCampus.patchCampus- Error - ${error.message}`);	
			return Promise.resolve(error);
		}
	} else {
		return util.buildResponse({error: true,
		message: 'patchCampus - required fields are undefined'});
	}
	
};