ReportComment.addComment = async (PINProjectID, PDFReportID, Comment, ReportActivity, PDFReportTemplateID) => { 
	try { 
		return await ReportComment.addComment(PDFReportID, Comment, ReportActivity, CreateBy);
		
	} catch (error) {
		return error;
	}
};